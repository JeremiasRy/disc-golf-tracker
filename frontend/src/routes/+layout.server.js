import { redirect } from '@sveltejs/kit';

/** @type {import("../../.svelte-kit/types/src/routes/$types").LayoutServerLoad}*/
export const load = async ({ cookies, url }) => {
    const token = cookies.get("auth_token")

    if (!token && url.pathname !== "/login") {
        throw redirect(302, '/login');
    }

    if (!token) {
        return {
            user: null
        };
    }

    // TODO: proper validation of session token
    const response = await fetch(`http://localhost:8800/users/${token}`)

    if (response.ok) {
        const user = await response.json();

        if (!url.pathname.startsWith('/app')) {
            throw redirect(302, '/app');
        }

        return { user }
    }

    return {
        user: null
    };
};

