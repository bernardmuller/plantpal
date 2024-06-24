import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { VitePWA } from "vite-plugin-pwa";

export default defineConfig({
	base: "/",
	preview: {
		port: 8000,
		strictPort: true,
	},
	server: {
		port: 8080,
		strictPort: true,
		host: true,
		origin: "http://0.0.0.0:8000",
	},
	plugins: [
		react(),
		VitePWA({
			registerType: "prompt",
			devOptions: {
				enabled: true,
			},
			injectRegister: "auto",
			includeAssets: [
				"favicon.ico",
				"apple-touch-icon.png",
				"maskable_icon.png",
			],
			manifest: {
				name: "PlantPal",
				short_name: "PlantPal",
				description: "PlantPal - Your plant companion app",
				theme_color: "#171717",
				background_color: "#f0e7db",
				display: "standalone",
				scope: "/",
				start_url: "/",
				orientation: "portrait",
				icons: [
					{
						src: "/android-chrome-192x192.png",
						sizes: "192x192",
						type: "image/png",
						purpose: "favicon",
					},
					{
						src: "/android-chrome-512x512.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "favicon",
					},
					{
						src: "/apple-touch-icon.png",
						sizes: "180x180",
						type: "image/png",
						purpose: "apple touch icon",
					},
					{
						src: "/maskable_icon.png",
						sizes: "512x512",
						type: "image/png",
						purpose: "any maskable",
					},
				],
			},
		}),
	],
});
