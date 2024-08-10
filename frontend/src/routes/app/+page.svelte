<script>
    import { useAuth } from "$lib/auth";
    import Navigation from "$lib/Navigation.svelte";
    const userStore = useAuth();
    /**@type {import("../../lib/types").User | null}*/
    let user;

    userStore.subscribe((value) => {
        user = value;
    });
</script>

{#if user}
    <h1>Welcome {user.name}</h1>
{/if}
