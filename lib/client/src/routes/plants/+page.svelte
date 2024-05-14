<!--export const prerender = true;-->
<script>
    import {createQuery, useQueryClient} from '@tanstack/svelte-query'
    import {Skeleton} from "$lib/components/ui/skeleton";
    import {getAllPlants} from "$lib/api/getAllPlants";

    export let data
    // const queryClient = useQueryClient()
    const plantsQuery = createQuery({
        queryKey: ['plants'],
        queryFn: () => getAllPlants(),
        initialData: data.plants || [],
    })

    const currentUser = data.currentUser

</script>


<div class="w-full flex justify-between py-4">
    <h1>Your Plants</h1>
</div>

<div>
    {#if $plantsQuery.isFetching}
        <div class="flex flex-col gap-4">
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
            <Skeleton class="w-[350px] h-8"/>
        </div>
    {:else if $plantsQuery.isError}
        <p>Error: {$plantsQuery.error.message}</p>
    {:else if $plantsQuery.isSuccess}
        {#if $plantsQuery.data?.data?.length === 0 || $plantsQuery.data.data === undefined}
            <p>No plants found</p>
        {:else}
            {#each $plantsQuery.data?.data as plant}
                <div class="flex gap-4">
                    <p class="font-semibold">{plant.ID}</p>
                    <p>{plant.Common}</p>
                </div>
            {/each}
        {/if}
    {/if}
</div>
