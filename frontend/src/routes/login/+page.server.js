import { redirect } from "@sveltejs/kit";

/** @type {import("./$types").Actions} */
export const actions = {
    login: async ({ request, cookies }) => {
        const data = await request.formData();
        const userId = data.get('userid');

        if (!userId) {
            return { error: "Provide something" }
        }

        const response = await fetch(`http://localhost:8800/users/${userId}`)
        console.log(response)
        if (response.ok) {
            const user = await response.json();
            cookies.set("auth_token", user.ID, { path: "/" })
            throw redirect(302, "/")
        }
        cookies.delete("auth_token", { path: "/" })
        return { error: "Failed to login" }
    }
};