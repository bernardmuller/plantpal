import { Plant } from "../../../lib/http/client/getAllPlants";
import { useLoaderData } from "react-router-dom";

type PageContainerProps = {
	children: React.ReactNode;
};

const PageContainer = ({ children }: PageContainerProps) => (
	<div className="w-screen flex flex-col items-center">
		<div className="w-full md:w-9/12 lg:w-7/12 p-4">{children}</div>
	</div>
);

export default function PlantDetailPage() {
	const { plant } = useLoaderData() as { plant: { plant: Plant } };
	if (!plant) return <div>Loading...</div>;
	return <PageContainer>Plant Detail</PageContainer>;
}
