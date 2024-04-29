<script>
    import {createQuery} from '@tanstack/svelte-query'
    import {fetchTodos} from "../lib/api/fetchTodos";

    const todosQuery = createQuery({
        queryKey: ['todos'],
        queryFn: () => fetchTodos(),
    })

    // @ts-ignore: data will exist, not sure why it's not being picked up
    console.log(todosQuery.data)

</script>


<h1>Welcome to SvelteKit</h1>
<div>
    {#if $todosQuery.isLoading}
        <p>Loading...</p>
    {:else if $todosQuery.isError}
        <p>Error: {$todosQuery.error.message}</p>
    {:else if $todosQuery.isSuccess}
        {#each $todosQuery.data as todo}
            <div class="flex gap-4">
                <p class="font-semibold">{todo.id}</p>
                <p>{todo.title}</p>
                <p class={todo.completed ? `text-green-300` : "text-stone-400" }>{todo.completed ? "Done" : "To Do"}</p>
            </div>
        {/each}
    {/if}
</div>
