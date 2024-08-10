/** @type {import("./$types").PageServerLoad} */
export const load = async ({ fetch }) => {
    const response = await fetch("http://localhost:8800/courses")
    if (response.ok) {
        /**@type {import("../../../lib/types").Course[]} */
        const courses = await response.json()
        return { courses }
    }

    return { courses: [] }
}