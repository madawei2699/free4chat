import Config

# We don't run a server during test. If one is required,
# you can enable the server option below.
config :free4chat, Free4chatWeb.Endpoint,
  http: [ip: {127, 0, 0, 1}, port: 4002],
  secret_key_base: "0TBV4M2xdbgLSj9jm1bkMLfjVkMGWDHrcKik/NQ67d2KwLc30fN4qJ13DHJtbbU5",
  server: false

# Print only warnings and errors during test
config :logger, level: :warn

# Initialize plugs at runtime for faster test compilation
config :phoenix, :plug_init_mode, :runtime
