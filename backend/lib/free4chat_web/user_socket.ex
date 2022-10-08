defmodule Free4chatWeb.UserSocket do
  use Phoenix.Socket

  channel("room:*", Free4chatWeb.PeerChannel)

  @impl true
  def connect(_params, socket, _connect_info) do
    {:ok, socket}
  end

  @impl true
  def id(_socket), do: nil
end
