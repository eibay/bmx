import { createRouter, createWebHistory } from 'vue-router';
import HomeView from './views/HomeView.vue';
import CommandsView from './views/Commands.vue';


const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
  },
  {
    path: '/com',
    name: 'Commands',
    component: CommandsView,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
