defmodule Day1 do
  {:ok, file} = File.read("./input.txt")

  file = String.split(file, "\n\n")

  sum_map = fn x ->
    String.split(x, "\n")
    |> Enum.filter(fn x -> x != "" end)
    |> Enum.map(fn x -> String.to_integer(x) end)
    |> Enum.sum()
  end

  r =
    Enum.map(file, sum_map)
    |> Enum.max()

  IO.inspect(r)
end
