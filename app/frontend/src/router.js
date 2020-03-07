import Vue from 'vue';
import VueRouter from 'vue-router';

import SplashPage from './pages/SplashPage';
import MainPage from './pages/MainPage';

Vue.use(VueRouter);

const routes = [
    { path: '/', component: MainPage },
    { path: '/splash', component: SplashPage }
];

export const router = new VueRouter({mode: "abstract", routes});
