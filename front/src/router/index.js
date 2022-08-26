import {
	createRouter,
	createWebHashHistory,

} from 'vue-router'
import Login from '../components/login.vue'
import Home from '../components/home.vue'
import Welcome from '../components/welcome.vue'
import Vessel from '../components/pages/vessel.vue'
import view from '../components/pages/view.vue'
import Newstockin from '../components/pages/Newstockin.vue'
import Newstockout from '../components/pages/Newstockout.vue'
const routes = [
	{
		path: '/login',
		component: Login,
		meta: {
			title: "登录"
		},
	},
	{
		path: '/home',
		component: Home,
		meta: {
			title: "港口堆存费管理系统"
		},
		children: [{
			path: '/welcome',
			component: Welcome,

		},
		{
			path: '/vessel',
			component: Vessel,
		},
		{
			path:'/view',
			component: view,
		},
		{
			path:'/Newstockin',
			component: Newstockin,
		},
		{
			path:'/Newstockout',
			component: Newstockout,
		},
		],
	},
]

const router = createRouter({
	history: createWebHashHistory(),
	routes
})


router.beforeEach((to, from, next) => {
	if (to.meta.title) {
		document.title = to.meta.title;
	}
	if (to.path == '/login') return next();

	const flagStr = window.sessionStorage.getItem("username");
	if (!flagStr) return next('/login');
	next();
})

export default router
