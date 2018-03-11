defmodule Gateway.Zombie.Api do
  @moduledoc """
  Implimentation for the Zombie Service API.
  """

  @behaviour Gateway.Zombie

  @user_agent [{"User-agent", "Gateway/0.42"}]

  def status(driver_id) do
    HTTPoison.get(url(driver_id), @user_agent) |> handle_response
  end

  def url(driver_id) do
    host() <> ":" <> port() <> "/drivers/#{driver_id}"
  end

  def handle_response({:ok, %{status_code: 200, body: body}}) do
    {:ok, parse(body)}
  end

  def handle_response({_, %{body: body}}) do
    {:error, parse(body)}
  end

  def parse(body) do
    Poison.Parser.parse!(body)
  end

  defp host do
    Application.get_env(:gateway, :zombie_host)
  end

  defp port do
    Application.get_env(:gateway, :zombie_port)
  end
end
