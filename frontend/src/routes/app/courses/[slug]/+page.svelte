<script>
    /** @type {import('./$types').PageData} */
    export let data;

    const { course } = data;
</script>

<h1>{course.name}</h1>

<form method="POST" action="?/startNewRound">
    <input type="hidden" name="courseId" value={course.ID} />
    <button type="submit">Start a new round?</button>
</form>
