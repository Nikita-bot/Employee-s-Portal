<template>
  <div class="sideMenu">
    <div 
      v-for="item in menuItems" 
      :key="item.name" 
      class="menuEl"
      @click="selectMenuItem(item)"
      :class="{ active: selectedItem === item.name}"
    >
      {{ item.title }}
    </div>
  </div>
</template>

<script setup>
  import { ref, computed, watch } from 'vue';
  import { useRoute, useRouter } from 'vue-router';

  const route = useRoute();
  const router = useRouter();

  const menuItems = ref([
    { name: 'tasks', title: 'Задачи', route: '/tasks' },
    { name: 'edo', title: 'ЭДО', route: '/edo' },
    { name: 'ecp', title: 'ЭЦП', route: '/ecp' },
    { name: 'analytics', title: 'Аналитика', route: '/analytics' },
    { name: 'calendar', title: 'Календарь', route: '/calendar' },
    { name: 'knowledge', title: 'База знаний', route: '/knowledge' }
  ]);

  const selectedItem = ref('');

  // Определяем, находимся ли мы на странице профиля
  const isProfilePage = computed(() => route.path === '/user');

  // Сбрасываем активный элемент при переходе на страницу пользователя
  watch(() => route.path, (newPath) => {
    if (newPath === '/user') {
      selectedItem.value = '';
    }
  });

  const selectMenuItem = (item) => {
    selectedItem.value = item.name;
    router.push(item.route);
  };
</script>

<style scoped>

        .sideMenu{
            background-color: #5662DE;
            width: 250px;
            height: 100%;
            text-align: center;

        }

        .menuEl{
            justify-content: center;
            color: #FFFFFF;
            font-size: 25px;
            height: 50px;
            padding-top: 5px;
            border-radius: 15px;
            margin-top: 20px;
        }

        .menuEl:hover{
            background-color: #020f9e;
        }

        .menuEl.active {
            background-color: #020f9e;
        }

</style>