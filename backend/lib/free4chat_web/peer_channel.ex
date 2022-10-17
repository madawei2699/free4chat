defmodule Free4chatWeb.PeerChannel do
  @moduledoc """
  The main channel used to communicate with a peer,
  i.e. exchange media events and any additional messages
  """
  use Phoenix.Channel

  require Logger

  @impl true
  def join("room:" <> room_id, params, socket) do
    simulcast? = Map.get(params, "isSimulcastOn")

    case :global.whereis_name(room_id) do
      :undefined ->
        Free4chat.Room.start(%{room_id: room_id, simulcast?: simulcast?}, name: {:global, room_id})

      pid ->
        {:ok, pid}
    end
    |> case do
      {:ok, room_pid} ->
        do_join(socket, room_pid, room_id)

      {:error, {:already_started, room_pid}} ->
        do_join(socket, room_pid, room_id)

      {:error, reason} ->
        Logger.error("""
        Failed to start room.
        Room: #{inspect(room_id)}
        Reason: #{inspect(reason)}
        """)

        {:error, %{reason: "failed to start room"}}
    end
  end

  defp do_join(socket, room_pid, room_id) do
    peer_id = "#{UUID.uuid4()}"
    # TODO handle crash of room?
    Process.monitor(room_pid)
    send(room_pid, {:add_peer_channel, self(), peer_id})

    {:ok,
     Phoenix.Socket.assign(socket, %{room_id: room_id, room_pid: room_pid, peer_id: peer_id})}
  end

  @impl true
  def handle_in("mediaEvent", %{"data" => event}, socket) do
    send(socket.assigns.room_pid, {:media_event, socket.assigns.peer_id, event})

    {:noreply, socket}
  end

  @impl true
  def handle_in("textEvent", %{"data" => event}, socket) do
    send(socket.assigns.room_pid, {:text_event, socket.assigns.peer_id, event})

    {:noreply, socket}
  end

  @impl true
  def handle_info(
        {:DOWN, _ref, :process, _pid, _reason},
        socket
      ) do
    {:stop, :normal, socket}
  end

  @impl true
  def handle_info({:media_event, event}, socket) do
    push(socket, "mediaEvent", %{data: event})

    {:noreply, socket}
  end

  @impl true
  def handle_info({:text_event, event}, socket) do
    push(socket, "textEvent", %{data: event})

    {:noreply, socket}
  end

  @impl true
  def handle_info({:simulcast_config, simulcast_config}, socket) do
    push(socket, "simulcastConfig", %{data: simulcast_config})

    {:noreply, socket}
  end
end
