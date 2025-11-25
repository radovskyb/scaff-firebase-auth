<template>
    <h1>Home</h1>

    <p>Hi {{ authStore.user?.displayName ?? 'there!' }}</p>
    <div v-if="authStore.isAuthenticated">
        <button @click="signOut">
            Sign out
        </button>
    </div>

    <div v-else>
        <button @click="signinWithGoogle">
            Sign in with Google
        </button>
    </div>

    <button @click="sendRequestToAPI">Call API</button>
</template>

<script setup>
import {useAuthStore} from '@/stores/authStore';
import {protectedRoute} from '@/services/apiService';
import {useRouter, useRoute} from 'vue-router';
import {watch} from 'vue';

const authStore = useAuthStore();
const router = useRouter();
const route = useRoute();

const sendRequestToAPI = async () => {
    try {
        const resp = await protectedRoute();
        console.log(resp);
    } catch (error) {
        console.error(error);
    }
};

const signinWithGoogle = async () => {
    try {
        await authStore.signInWithGoogle();
        router.push('/');
    } catch (error) {
        console.error(error);
    }
};

const signOut = async () => {
    try {
        await authStore.logout();
        router.push('/');
    } catch (error) {
        console.error(error);
    }
};

watch(() => authStore.isAuthenticated, (isAuthenticated) => {
    // Watch for authentication changes here if needed.
    if (isAuthenticated) {
        console.log('Signed in');
    } else {
        console.log('Signed out');
    }
});
</script>
