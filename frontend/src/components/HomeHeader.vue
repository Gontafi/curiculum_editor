<!-- eslint-disable -->
<template>
    <div id="header">
        <header>
            <div class="header-title-container">
                <h1 class="header-title">Создания учебного плана</h1>
            </div>
            <div class="header-buttons">
                <router-link to="/login" class="header-button">Зайти</router-link>
                <router-link to="/signup" class="header-button">Зарегистрироваться</router-link>
            </div>
        </header>
    </div>
</template>
<!-- eslint-disable -->
<script>
export default {
    name: "HomeHeader",
};
</script>
<!-- eslint-disable -->
<style>
header {
    background-color: #3c66aa;
    /* Blue-light color */
    color: white;
    padding: 5px;
    display: flex;
    justify-content: space-between;
    /* Align title and buttons horizontally */
    align-items: center;
    /* Align title and buttons vertically */
}

.header-title-container {
    flex-grow: 1;
    /* Allow the title to take up available space */
}

.header-title {
    font-size: 20px;
    /* Adjust the font size to make it smaller */
    margin-right: auto;
    /* Push the title to the left */
    margin-left: 20px;
    /* Add left margin to separate it from the edge */
}

.header-buttons {
    position: absolute;
    top: px;
    right: 20px;
}

/* Style for header button links */
.header-button {
    padding: 8px 8px;
    color: #fff;
    background-color: #3c66aa;
    border: none;
    border-radius: 4px;
    text-decoration: none;
    font-weight: bold;
    transition: background-color 0.2s;
}

.header-button:hover {
    background-color: #2c3e50;
    text-decoration: none;
}

.header-button + .header-button {
    margin-left: 10px;
}
</style>