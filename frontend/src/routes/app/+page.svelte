<script>
    import { useAuthContext } from "$lib";
    const userStore = useAuthContext();
    /**@type {import("../../lib/types").User | null}*/
    let user;

    userStore.subscribe((value) => {
        user = value;
    });
</script>

{#if user}
    <h1>Welcome {user.name}</h1>
{/if}
