import { Plant } from "../../lib/http/client/getAllPlants";
import { Link, useLoaderData, useNavigate } from "react-router-dom";
import {
	Card,
	CardContent,
	// CardHeader,
	// CardTitle,
} from "../../components/ui/card";
// import {
// 	DropdownMenu,
// 	DropdownMenuContent,
// 	DropdownMenuItem,
// 	DropdownMenuLabel,
// 	DropdownMenuSeparator,
// 	DropdownMenuTrigger,
// } from "../../components/ui/dropdown-menu";
import { BookmarkIcon, CalendarIcon, HomeIcon } from "@radix-ui/react-icons";

export default function Home() {
	const { plants } = useLoaderData() as { plants: { plants: Plant[] } };
	const navigate = useNavigate();
	const handleClick = (plantId: string) => navigate(`/plants/${plantId}`);
	if (!plants) return <div>Loading...</div>;
	return (
		<>
			<div className="flex min-h-screen w-full">
				<div className="hidden h-full border-r bg-muted/40 lg:block">
					<div className="flex h-full max-h-screen flex-col gap-2">
						<div className="flex h-[60px] items-center border-b px-6">
							<Link
								to="#"
								className="flex items-center gap-2 font-semibold"
							>
								{/* < className="h-6 w-6" /> */}
								<span>PlantPal</span>
							</Link>
						</div>
						<div className="flex-1 overflow-auto py-2 h-full">
							<nav className="grid items-start px-4 text-sm font-medium ">
								<Link
									to="#"
									className="flex items-center gap-3 rounded-lg px-3 py-2 text-primary transition-all hover:bg-muted"
								>
									<HomeIcon className="h-4 w-4" />
									My Plants
								</Link>
								<Link
									to="#"
									className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:bg-muted"
								>
									<CalendarIcon className="h-4 w-4" />
									Watering Schedule
								</Link>
								<Link
									to="#"
									className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:bg-muted"
								>
									<BookmarkIcon className="h-4 w-4" />
									Plant Care Tips
								</Link>
								<Link
									to="#"
									className="flex items-center gap-3 rounded-lg px-3 py-2 text-muted-foreground transition-all hover:bg-muted"
								>
									{/* <CogIcon className="h-4 w-4" /> */}
									Settings
								</Link>
							</nav>
						</div>
					</div>
				</div>
				<div className="flex flex-col w-full">
					<header className="flex h-14 lg:h-[60px] items-center gap-4 border-b bg-muted/40 px-6">
						<Link to="#" className="lg:hidden">
							{/* <LeafIcon className="h-6 w-6" /> */}
							<span className="sr-only">Home</span>
						</Link>
						<div className="flex-1">
							<h1 className="text-lg font-semibold">My Plants</h1>
						</div>
						{/* <DropdownMenu>
							<DropdownMenuTrigger asChild>
								<Button
									variant="ghost"
									size="icon"
									className="rounded-full border w-8 h-8"
								>
									<img
										src="/placeholder.svg"
										width="32"
										height="32"
										className="rounded-full"
										alt="Avatar"
									/>
									<span className="sr-only">
										Toggle user menu
									</span>
								</Button>
							</DropdownMenuTrigger>
							<DropdownMenuContent align="end">
								<DropdownMenuLabel>
									My Account
								</DropdownMenuLabel>
								<DropdownMenuSeparator />
								<DropdownMenuItem>Settings</DropdownMenuItem>
								<DropdownMenuItem>Support</DropdownMenuItem>
								<DropdownMenuSeparator />
								<DropdownMenuItem>Logout</DropdownMenuItem>
							</DropdownMenuContent>
						</DropdownMenu> */}
					</header>
					<main className="flex-1 flex flex-col gap-4 p-4 md:gap-8 md:p-6">
						<div className="grid gap-4 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4">
							{plants?.plants?.map((plant: Plant) => (
								<Card onClick={() => handleClick(plant.ID)}>
									<CardContent className="grid gap-2 pt-6">
										<div className="flex flex-col gap-1">
											<div className="font-medium">
												{plant.Common}
											</div>
											<div className="text-sm text-muted-foreground">
												{plant.Family}
											</div>
										</div>
									</CardContent>
								</Card>
							))}
						</div>
					</main>
				</div>
			</div>
		</>
	);
}
