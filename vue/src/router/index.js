import { createRouter, createWebHistory } from 'vue-router'
import Index from '@/views/Index.vue'
import Types from '@/views/Types.vue'
import Relayers from '@/views/Relayers.vue'
import Explorer from '@/views/Explorer.vue'

const routerHistory = createWebHistory()
const routes = [
	{
		path: '/',
		component: Index
	},
	{ path: '/types', component: Types },
	{ path: '/relayers', component: Relayers },
	{ path: '/explorer', component: Explorer }
]

const router = createRouter({
	history: routerHistory,
	routes
})

export default router
