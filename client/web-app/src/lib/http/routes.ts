const getBaseURL = () => {
	if (process.env.NODE_ENV === "development") {
		return "http://localhost:8001";
	} else {
		return "https://plantpal-api.bernardmuller.co.za";
	}
};

const apiRoutes = {
	getAllPlants: `${getBaseURL}/plants`,
	getPlantById: (id: string) => `${getBaseURL}/plants/${id}`,
};
export default apiRoutes;
