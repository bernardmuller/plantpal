import { Plant } from "../../../lib/http/client/getAllPlants";
import { useLoaderData, useNavigate } from "react-router-dom";
import { ArrowLeftIcon } from "@radix-ui/react-icons";
import {
	Card,
	CardContent,
	CardHeader,
	CardTitle,
} from "../../../components/ui/card";
import { Button } from "../../../components/ui/button";

type PageContainerProps = {
	children: React.ReactNode;
};

const PageContainer = ({ children }: PageContainerProps) => (
	<div className="w-screen flex flex-col items-center">
		<div className="w-full md:w-9/12 lg:w-7/12 p-4">{children}</div>
	</div>
);

export default function PlantDetailPage() {
	const navigate = useNavigate();
	const { plant } = useLoaderData() as { plant: Plant };
	if (!plant) return <div>Loading...</div>;
	return (
		<PageContainer>
			<Button
				variant="link"
				className="px-4 py-2"
				onClick={() => navigate("/")}
			>
				<ArrowLeftIcon height={30} width={30} />
			</Button>
			<div className="grid md:grid-cols-2 gap-6 lg:gap-12 items-start max-w-6xl px-4 mx-auto py-6 pt-2">
				<div className="grid gap-4 md:gap-10 items-start">
					<div className="flex items-center justify-center">
						{/* <img
							src="/placeholder.svg"
							alt="Plant Image"
							width={600}
							height={600}
							className="rounded-lg object-cover w-full aspect-square"
						/> */}
					</div>
					<div className="grid gap-2">
						<h1 className="text-2xl font-bold">{plant.Common}</h1>
						<div className="text-muted-foreground">
							<span className="font-medium">Species:</span>{" "}
							{plant.Family}
						</div>
					</div>
				</div>
				<div className="grid gap-6">
					<Card>
						<CardHeader>
							<CardTitle>Care Instructions</CardTitle>
						</CardHeader>
						<CardContent className="grid gap-4 text-sm">
							<div>
								<div className="font-medium">Water</div>
								<p>
									Water your Monstera Deliciosa when the top
									inch of soil is dry. Avoid letting the soil
									become completely dry.
								</p>
							</div>
							<div>
								<div className="font-medium">Sunlight</div>
								<p>
									Bright, indirect light is best. Avoid direct
									sunlight, which can scorch the leaves.
								</p>
							</div>
							<div>
								<div className="font-medium">Humidity</div>
								<p>
									Mist the leaves regularly or use a pebble
									tray to increase humidity around the plant.
								</p>
							</div>
							<div>
								<div className="font-medium">Fertilizer</div>
								<p>
									Feed your Monstera Deliciosa with a balanced
									liquid fertilizer every 2-3 months during
									the growing season.
								</p>
							</div>
						</CardContent>
					</Card>
					<Card>
						<CardHeader>
							<CardTitle>Notes</CardTitle>
						</CardHeader>
						<CardContent className="grid gap-4 text-sm">
							<div>
								<div className="font-medium">
									Repotting Reminder
								</div>
								<p>
									Remember to repot the Monstera Deliciosa in
									the spring every 2-3 years.
								</p>
							</div>
							<div>
								<div className="font-medium">Leaf Cleaning</div>
								<p>
									Wipe down the leaves with a damp cloth every
									few weeks to keep them clean and healthy.
								</p>
							</div>
							<div>
								<div className="font-medium">Propagation</div>
								<p>
									Take a stem cutting in the spring and
									propagate a new Monstera Deliciosa plant.
								</p>
							</div>
						</CardContent>
					</Card>
				</div>
			</div>
		</PageContainer>
	);
}
