defmodule Free4chatWeb.RoomController do
  use Free4chatWeb, :controller

  def scrape(conn, %{"room_id" => id}) do
    response =
      Membrane.TelemetryMetrics.Reporter.scrape(Free4chatReporter)
      |> Map.get({:room_id, id})
      |> inspect(pretty: true, limit: :infinity)

    text(conn, response)
  end
end
