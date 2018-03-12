use Mix.Config

config :gateway,
  port: 3000,
  nsqd_topic: "driver",
  nsqd_host: "127.0.0.1:4150",
  zombie_host: "127.0.0.1:3002"
