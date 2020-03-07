import './assets/css/reset.css';
import "./assets/css/main.css";

import Vue from 'vue';
import App from './App.vue';

import { router } from './router';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Wails.Init(() => {
	new Vue({
		router, render: h => h(App), mounted() {}
	}).$mount('#app');
	router.replace('/splash');
	setTimeout(() => { router.replace('/') }, 2000);
});
