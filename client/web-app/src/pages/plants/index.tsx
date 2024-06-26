import { Plant } from "../../lib/http/client/getAllPlants";
import { useLoaderData, useNavigate } from "react-router-dom";
import {
	Table,
	TableBody,
	TableCaption,
	TableCell,
	TableHead,
	TableHeader,
	TableRow,
} from "../../components/ui/table";

type PageContainerProps = {
	children: React.ReactNode;
};

const PageContainer = ({ children }: PageContainerProps) => (
	<div className="w-screen flex flex-col items-center">
		<div className="w-full md:w-9/12 lg:w-7/12 p-4">{children}</div>
	</div>
);

export default function Home() {
	const { plants } = useLoaderData() as { plants: { plants: Plant[] } };
	const navigate = useNavigate();
	const handleClick = (plantId: string) => navigate(`/plants/${plantId}`);
	if (!plants) return <div>Loading...</div>;
	return (
		<PageContainer>
			<Table>
				<TableCaption>A list of your plants</TableCaption>
				<TableHeader>
					<TableRow>
						<TableHead>Name</TableHead>
						<TableHead>Family</TableHead>
						<TableHead>Category</TableHead>
					</TableRow>
				</TableHeader>
				<TableBody>
					{plants?.plants?.map((plant: Plant) => (
						<TableRow
							key={plant.ID}
							onClick={() => handleClick(plant.ID)}
						>
							<TableCell className="font-medium w-1/2">
								{plant.Common}
							</TableCell>
							<TableCell>{plant.Family}</TableCell>
							<TableCell>{plant.Category}</TableCell>
						</TableRow>
					))}
				</TableBody>
			</Table>
		</PageContainer>
	);
}
