<script>
    import { useAuthContext } from "$lib";

    /** @type {import('./$types').PageData} */
    export let data;
    const userStore = useAuthContext();
    /**
     * @type {import("$lib/types").User | null}
     */
    let user;

    userStore.subscribe((value) => {
        user = value;
    });
    const { course } = data;
</script>

<h1>{course.name}</h1>

{#if user}
    <form method="POST" action="?/startNewRound">
        <input type="hidden" name="courseId" value={course.ID} />
        <input type="hidden" name="userId" value={user.ID} />
        <button type="submit">Create a new round?</button>
    </form>
{/if}
