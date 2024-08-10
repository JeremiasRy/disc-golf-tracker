<script>
    /** @type {import('./$types').ActionData} */
    export let form;
</script>

<form method="POST" action="?/login">
    <label
        >User Id
        <input type="text" id="userid" name="userid" required />
    </label>

    {#if form?.error}
        <p style="color: red;">{form.error}</p>
    {/if}

    <button type="submit">Login</button>
</form>
