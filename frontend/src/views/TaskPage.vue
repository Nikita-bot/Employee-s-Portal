<template>
    <div id="app">
        <HeaderComponent />
        <div class="container">
            <SideMenu />
            <TaskGrid />
        </div>
    </div>    
</template>

<script setup>
    import HeaderComponent from '@/components/HadderComponent.vue';
    import SideMenu from '@/components/SideMenu.vue';
    import TaskGrid from '@/components/TaskGrid.vue';
</script>