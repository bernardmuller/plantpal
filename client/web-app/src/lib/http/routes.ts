const apiRoutes = {
	getAllPlants: `http://localhost:8001/plants`,
	getPlantById: (id: string) => `http://localhost:8001/plants/${id}`,
};
export default apiRoutes;
