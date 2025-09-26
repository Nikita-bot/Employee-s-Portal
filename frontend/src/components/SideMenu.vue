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
  import { ref, watch } from 'vue';
  import { useRoute, useRouter } from 'vue-router';
  import { useUserStore } from '@/stores/user';

  const route = useRoute();
  const router = useRouter();
  const userStore = useUserStore();
  if (!userStore.userData.id) {
    router.push('/login')
  }


  const menuItems = ref(userStore.userData.employee.department == "отдел автоматизированных систем управления" ? 
    [
      { name: 'tasks', title: 'Задачи', route: '/tasks' },
      // { name: 'edo', title: 'ЭДО', route: '/edo' },
      // { name: 'ecp', title: 'ЭЦП', route: '/ecp' },
      // { name: 'analytics', title: 'Аналитика', route: '/analytics' },
      // { name: 'calendar', title: 'Календарь', route: '/calendar' },
      { name: 'portals', title: 'Порталы', route: '/potals' },
      { name: 'knowledge', title: 'База знаний', route: '/knowledge' }
    ]:
    [
      { name: 'tasks', title: 'Задачи', route: '/tasks' },
      // { name: 'edo', title: 'ЭДО', route: '/edo' },
      // { name: 'ecp', title: 'ЭЦП', route: '/ecp' },
      // { name: 'analytics', title: 'Аналитика', route: '/analytics' },
      // { name: 'calendar', title: 'Календарь', route: '/calendar' },
      { name: 'portals', title: 'Порталы', route: '/potals' },
    ]
  );

  const selectedItem = ref('');

 const updateSelectedItem = () => {
  if (route.path === '/user') {
    selectedItem.value = '';
    return;
  }

  const currentItem = menuItems.value.find(item => 
    route.path.startsWith(item.route)
  );
  
  selectedItem.value = currentItem ? currentItem.name : '';
};

 
  updateSelectedItem();

  
  watch(() => route.path, updateSelectedItem);

  const selectMenuItem = (item) => {
    if (item.name === 'knowledge') {
      window.location.href = 'http://documents.kkbsmp.ru/view/list';
    } else if (item.name === 'portals'){
      window.location.href = 'http://portals.kkbsmp.ru/';
    } else if (item.name == 'edo'){
      window.location.href = 'http://edo.bsmp42.ru/index.php'
    }
    else {
      router.push(item.route);
    }
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
            transition: background-color 0.2s;
        }

        .menuEl:hover{
            background-color: #020f9e;
        }

        .menuEl.active {
            background-color: #020f9e;
        }

</style>