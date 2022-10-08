defmodule Free4chatWeb.Router do
  use Free4chatWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", Free4chatWeb do
    pipe_through :api
  end

  # Enables LiveDashboard only for development
  #
  # If you want to use the LiveDashboard in production, you should put
  # it behind authentication and allow only admins to access it.
  # If your application does not have an admins-only section yet,
  # you can use Plug.BasicAuth to set up some basic authentication
  # as long as you are also using SSL (which you should anyway).
  pipeline :dashboard do
    import Plug.BasicAuth
    plug :basic_auth, username: System.fetch_env!("DASHBOARD_AUTH_USERNAME"), password: System.fetch_env!("DASHBOARD_AUTH_PASSWORD")
  end

  if Mix.env() in [:dev, :test] do
    import Phoenix.LiveDashboard.Router

    scope "/" do
      pipe_through [:fetch_session, :protect_from_forgery, :dashboard]

      live_dashboard "/dashboard", metrics: Free4chatWeb.Telemetry
    end
  end
end
