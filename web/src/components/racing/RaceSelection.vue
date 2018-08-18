<template>
    <tr :class="{
            'scratched': selection.is_scratched
        }">
        <th class="sel-number">
            {{selection.number}}
        </th>
        <td class="sel-img">
            <img :src="selection.image_url" class="image" alt="">
        </td>
        <td class = "sel-info">
            <div class="name">{{selection.name}} ({{selection.barrier}})</div>
            <div class="jockey"><span class="j-icon">J</span>{{selection.jockey}} {{selection.jockey_weight}}</div>
        </td>
        <td class="price fixed win">
            4.00
        </td>
        <td class="price fixed place">
            1.40
        </td>
        <td class="price tote win">
            3.20
        </td>
        <td class="price stote place">
            1.10
        </td>
    </tr>
</template>

<script>
export default {
  name: 'RaceSelection',
  props: ['selection'],
};
</script>

<style lang="scss" scoped>
tr{
    height: 70px;
}

.sel-info {
    width: 100%;
}

.sel-img {
    min-width: 65px;
}

.jockey {
    font-size: 0.7em;
}

.scratched td {
    text-decoration: underline;
}

.j-icon {
    font-family: 'Quicksand';
    font-size: 0.8em;
    display: inline-block;
    background-color: #04518D;
    color: white;
    width: 12px;
    text-align: center;
    border-radius: 100%;
    margin-right: 5px;
}

</style>
