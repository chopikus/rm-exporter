<script lang="ts">
    import { ToolbarButton, Checkbox, Tooltip } from "flowbite-svelte";
    import { ArrowUpOutline } from "flowbite-svelte-icons";

    let {
        id,
        path,
        onBack,
        isItemChecked,
        isItemIndeterminate,
        itemCheckUpdate,
    } = $props();
</script>

<nav class="bg-blue-50 text-blue-800 py-2.5 w-full sticky top-0 h-14">
    <div class="w-full h-full flex flex-row items-center">
        <div class="flex-1">
            <Checkbox
                indeterminate={isItemIndeterminate(id)}
                bind:checked={
                    () => isItemChecked(id), (v) => itemCheckUpdate(id, v)
                }
                class="w-4 h-4"
                style="margin-left: 57px;"
            />
        </div>
        <h1 class="font-bold mx-auto">Choose files to export</h1>
        <div class="flex-1">
            <div class="float-right mr-11">
                {#if path.length !== 0}
                    <ToolbarButton color="blue" name="Back" onclick={onBack}>
                        <ArrowUpOutline class="w-7 h-7" />
                    </ToolbarButton>
                {/if}
            </div>
        </div>
    </div>
</nav>
