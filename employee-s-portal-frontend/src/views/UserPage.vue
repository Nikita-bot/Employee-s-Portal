<template>
    <div id="app">
        <HeaderComponent />
        <div class="container">
            <SideMenu />
            <UserInfo/>
        </div>
    </div>    
</template>

<script setup>
    import HeaderComponent from '@/components/HadderComponent.vue';
    import SideMenu from '@/components/SideMenu.vue';
import UserInfo from '@/components/UserInfo.vue';
</script>