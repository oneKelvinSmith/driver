defmodule Gateway.Location.ProducerTest do
  use ExUnit.Case
  doctest Gateway

  alias Gateway.Location.Producer

  @tag :nsq
  test "exists" do
    driver_id = 42

    assert Producer.update_location(driver_id) == {:ok, "OK"}
  end
end
