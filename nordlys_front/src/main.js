import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import HomePage from "./pages/HomePage.vue";
import AboutPage from "./pages/AboutPage.vue";
import {createMemoryHistory, createRouter} from "vue-router";
import LoginPage from "./pages/LoginPage.vue";
import RegisterPage from "./pages/RegisterPage.vue";

const routes = [
    { path: '/', component: HomePage },
    { path: '/about', component: AboutPage },
    { path: '/login', component: LoginPage },
    { path: '/register', component: RegisterPage }
]

const router = createRouter({
    history: createMemoryHistory(),
    routes
})

createApp(App)
    .use(router)
    .mount('#app')
