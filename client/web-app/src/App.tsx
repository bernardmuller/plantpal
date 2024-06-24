import { useState } from "react";
import logo from "/android-chrome-512x512.png";
import "./App.css";

function App() {
	const [count, setCount] = useState(0);

	return (
		<>
			<div>
				<img src={logo} className="logo" alt="Plantpal logo" />
			</div>
			<h1>Plantpal</h1>
			<div className="card">
				<button onClick={() => setCount((count) => count + 1)}>
					count is {count}
				</button>
				<p>
					Edit <code>src/App.tsx</code> and save to test HMR
				</p>
			</div>
		</>
	);
}

export default App;
