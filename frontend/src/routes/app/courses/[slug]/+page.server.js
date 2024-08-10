import { error } from '@sveltejs/kit';

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