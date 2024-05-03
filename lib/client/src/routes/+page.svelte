<!--export const prerender = true;-->
<script>
    import {createQuery, useQueryClient} from '@tanstack/svelte-query'
    import {Button} from "$lib/components/ui/button";
    import {Skeleton} from "$lib/components/ui/skeleton";
    import {login} from "$lib/api/login";
    import {getAllPlants} from "$lib/api/getAllPlants";

    export let data
    const queryClient = useQueryClient()
    const plantsQuery = createQuery({
        queryKey: ['plants'],
        queryFn: () => getAllPlants(),
        initialData: data.plants ?? []
    })

    // const handleLogin = async () => {
    //     // await login().catch((err) => console.log("Error", err))
    // }
</script>


<div class="w-full flex justify-between">
    <h1>Welcome to PlantPal</h1>
    <a href="http://localhost:8080/auth/google?provider=google">
        Log in with Google
    </a>
</div>
