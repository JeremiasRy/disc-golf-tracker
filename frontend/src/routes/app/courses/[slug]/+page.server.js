import { error, redirect } from '@sveltejs/kit';

/** @type {import('./$types').PageServerLoad} */
export const load = async ({ params, fetch }) => {
    const courseId = params.slug;
    const response = await fetch(`http://localhost:8800/courses/${courseId}`);

    if (response.ok) {
        /**
         * @type {import("../../../../lib/types").Course}
         */
        const course = await response.json()
        return { course }
    }
    error(404, "Not found")
}

export const actions = {
    /**
    * Start a new round and redirect to the new round's page.
    * 
    * @param {import('@sveltejs/kit').RequestEvent} event - The request event object containing the request.
    * @returns {Promise<void>} Redirect to the new round
    */
    startNewRound: async ({ request }) => {
        const data = await request.formData()
        const courseId = Number(data.get('courseId'));
        const body = {
            course_id: courseId
        }

        const requestInit = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(body)
        }
        const response = await fetch("http://localhost:8800/rounds", requestInit)

        if (response.ok) {
            const round = await response.json()
            redirect(303, `/app/rounds/${round.ID}`)
        } else {
            throw error(response.status, 'Failed to create new round');
        }
    }
};