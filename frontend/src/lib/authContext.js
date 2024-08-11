import { writable } from 'svelte/store';
import { setContext, getContext } from 'svelte';
export const AUTH_CONTEXT = 'auth';

/** @type {import('svelte/store').Writable<import('$lib/types').User | null>} */
const userStore = writable(null);

/** @param {import('$lib/types').User | null} user */
export function setUser(user) {
    userStore.set(user);
}

export function initAuthContext() {
    setContext(AUTH_CONTEXT, userStore);
}

/**
 * @returns {import('svelte/store').Writable<import('$lib/types').User | null>}
 */
export function useAuthContext() {
    return getContext(AUTH_CONTEXT);
}