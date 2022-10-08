defmodule Free4chatWeb.PageController do
  use Free4chatWeb, :controller

  def healthcheck(conn, _params) do
    conn
    |> send_resp(200, "ok")
  end
end
