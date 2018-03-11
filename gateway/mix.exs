defmodule Gateway.MixProject do
  use Mix.Project

  def project do
    [
      app: :gateway,
      version: "0.1.0",
      elixir: "~> 1.6",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  def application do
    [
      extra_applications: [:logger],
      mod: {Gateway.Application, []}
    ]
  end

  defp deps do
    [
      {:cowboy, "~> 1.1"},
      {:httpoison, "~> 1.0"},
      {:plug, "~> 1.5"},
      {:poison, "~> 3.1"}
    ]
  end
end
