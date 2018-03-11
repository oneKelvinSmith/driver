use Mix.Config

config :gateway,
  port: "${PORT}",
  nsqd_topic: "${NSQD_TOPIC}",
  nsqd_host: "${NSQD_HOST}",
  nsqd_port: "${NSQD_PORT}",
  zombie_host: "${ZOMBIE_HOST}",
  zombie_port: "${ZOMBIE_PORT}"
