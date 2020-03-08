import './assets/css/reset.css';
import "./assets/css/main.css";
import 'metro4-dist/css/metro-icons.css';

import Vue from 'vue';
import App from './App.vue';

import AppButton from './components/AppButton';
import AppIcon from './components/AppIcon';

import { router } from './router';

Vue.config.productionTip = false;
Vue.config.devtools = true;

import * as Wails from '@wailsapp/runtime';

Vue.component('AppButton', AppButton);
Vue.component('AppIcon', AppIcon);

Wails.Init(() => {
	new Vue({
		router, render: h => h(App), mounted() {}
	}).$mount('#app');
	router.replace('/splash');
	setTimeout(() => { router.replace('/') }, 2000);
});
