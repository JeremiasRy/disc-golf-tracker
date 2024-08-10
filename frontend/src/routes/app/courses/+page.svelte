<script>
    /** @type {import('./$types').PageData} */
    export let data;

    const { courses } = data;
</script>

<h1>Courses</h1>
{#each courses as course}
    <a href={`/app/courses/${course.ID}`}>{course.name}</a>
{/each}
