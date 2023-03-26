const raw_input = await Deno.readTextFile("input.txt");

const input = raw_input.trim()
  .split("\n\n")
  .map((v) => v.split("\n").reduce((pv, cv) => pv + parseInt(cv), 0))
  .sort((a, b) => {
    if (a < b) return -1;
    if (a > b) return 1;
    return 0;
  })
  .pop();

console.log(input);
