use Mix.Config

config :gateway,
  port: "${PORT}",
  nsqd_topic: "${NSQD_TOPIC}",
  nsqd_host: "${NSQD_HOST}",
  zombie_host: "${ZOMBIE_HOST}"
