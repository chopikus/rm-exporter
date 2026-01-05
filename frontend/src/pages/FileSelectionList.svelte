<script lang="ts">
    import { Listgroup, Checkbox, P } from "flowbite-svelte";
    import { FolderSolid, FileLinesSolid } from "flowbite-svelte-icons";
    import { backend } from "../../wailsjs/go/models";
    type DocInfo = backend.DocInfo;

    let {
        items,
        isItemChecked,
        isItemIndeterminate,
        itemCheckUpdate,
        onItemClick,
    } = $props();
</script>

{#if items.length > 0}
    <Listgroup {items} let:item active={false}>
        <div class="flex flex-row justify-start items-center">
            <Checkbox
                bind:checked={
                    () => isItemChecked(item.Id),
                    (v) => itemCheckUpdate(item.Id, v)
                }
                indeterminate={isItemIndeterminate(item.Id)}
                class="mr-4 w-4 h-4"
            />

            <div
                class="flex flex-row justify-start items-center w-full hover:bg-gray-100"
                onclick={() => onItemClick(item)}
            >
                {#if item.IsFolder}
                    <FolderSolid class="mr-0.5" size="lg" />
                {:else}
                    <FileLinesSolid class="mr-0.5" size="lg" />
                {/if}
                <P size="xl">{item.Name}</P>
            </div>
        </div>
    </Listgroup>
{/if}
