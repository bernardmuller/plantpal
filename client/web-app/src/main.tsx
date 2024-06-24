import { Container, createRoot } from "react-dom/client";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./index.css";
import Home from "./pages/home";
import homeLoader from "./lib/http/loaders/home";
import PlantDetailPage from "./pages/plants/detail";
import plantDetailLoader from "./lib/http/loaders/plants-detail";

const router = createBrowserRouter([
	{
		path: "/",
		element: <Home />,
		loader: homeLoader,
	},
	{
		path: "/plants/:plantId",
		element: <PlantDetailPage />,
		loader: plantDetailLoader,
	},
]);

createRoot(document.getElementById("root") as Container).render(
	<RouterProvider router={router} />
);
