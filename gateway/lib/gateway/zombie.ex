defmodule Gateway.Zombie do
  @moduledoc """
  Specification for the Zombie Service api.
  """

  @typep driver_id :: integer

  @doc """
  Fetches zombie status for a given `driver_id`.
  """
  @callback status(driver_id :: driver_id) :: {:ok, term} | {:error, term}
end
