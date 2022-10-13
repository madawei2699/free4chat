defmodule Free4chat.TurnServer do
  @moduledoc false

  use GenServer

  alias Membrane.ICE.TURNManager

  @mix_env Mix.env()

  def start_link(_opts) do
    GenServer.start_link(__MODULE__, %{})
  end

  @impl true
  def init(state) do
    start_turn()
    {:ok, state}
  end

  defp start_turn() do
    turn_cert_file =
      case Application.fetch_env(:free4chat, :integrated_turn_cert_pkey) do
        {:ok, val} -> val
        :error -> nil
      end

    turn_mock_ip = Application.fetch_env!(:free4chat, :integrated_turn_ip)
    turn_ip = if @mix_env == :prod, do: {0, 0, 0, 0}, else: turn_mock_ip

    integrated_turn_options = [
      ip: turn_ip,
      mock_ip: turn_mock_ip,
      ports_range: Application.fetch_env!(:free4chat, :integrated_turn_port_range),
      cert_file: turn_cert_file
    ]

    tcp_turn_port = Application.get_env(:free4chat, :integrated_tcp_turn_port)
    TURNManager.ensure_tcp_turn_launched(integrated_turn_options, port: tcp_turn_port)

    if turn_cert_file do
      tls_turn_port = Application.get_env(:free4chat, :integrated_tls_turn_port)
      TURNManager.ensure_tls_turn_launched(integrated_turn_options, port: tls_turn_port)
    end
  end

end
