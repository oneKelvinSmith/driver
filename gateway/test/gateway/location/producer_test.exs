defmodule Gateway.Location.ProducerTest do
  use ExUnit.Case
  doctest Gateway

  alias Gateway.Location.Producer

  test "exists" do
    driver_id = 42

    assert Producer.update_location(driver_id) == {:ok, {:updated, driver_id}}
  end
end
