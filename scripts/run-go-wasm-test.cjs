const [runner, wasm] = process.argv.slice(2);

if (!runner || !wasm) {
  console.error("Usage: node run-go-wasm-test.cjs <wasm_exec_node.js> <test.wasm>");
  process.exit(2);
}

const safePath = process.env.PATH;
for (const key of Object.keys(process.env)) delete process.env[key];
if (safePath) process.env.PATH = safePath;

process.argv = [process.execPath, runner, wasm];
require(runner);
