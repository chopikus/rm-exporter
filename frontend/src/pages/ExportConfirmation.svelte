<script lang="ts">
    import {
        Button,
        ButtonGroup,
        Listgroup,
        Navbar,
        P,
        ToolbarButton,
    } from "flowbite-svelte";
    import {
        GetCheckedFiles,
        DirectoryDialog,
        SetExportOptions,
    } from "../../wailsjs/go/main/App.js";
    import {
        ArrowLeftOutline,
        CheckOutline,
        FileLinesSolid,
    } from "flowbite-svelte-icons";
    import { backend } from "../../wailsjs/go/models.js";
    import { push } from "svelte-spa-router";
    type DocInfo = backend.DocInfo;

    let pdf = $state(true);
    let rmdoc = $state(true);
    let location = $state("");
    let items: DocInfo[] = $state([]);

    GetCheckedFiles().then((result: DocInfo[]) => {
        items = result;
    });

    const selectDirectory = () => {
        DirectoryDialog().then((result: string) => {
            location = result;
        });
    };

    const onProceed = () => {
        SetExportOptions({ pdf, rmdoc, location }).then(() => {
            push("/export");
        });
    };

    const onBack = () => {
        push("/files");
    };
</script>

<div style="height: fit-content;">
    <Navbar color="blue" class="sticky top-0">
        <div>
            <ToolbarButton color="blue" name="Back" onclick={onBack}>
                <ArrowLeftOutline class="w-7 h-7" /></ToolbarButton
            >
        </div>
        <h1 class="font-bold m-auto">Export options</h1>
        <div class="invisible">
            <ToolbarButton color="blue" name="Back" onclick={onBack}>
                <ArrowLeftOutline class="w-7 h-7" /></ToolbarButton
            >
        </div>
    </Navbar>

    <main class="pr-7 pl-7 mt-3 w-full">
        <div class="flex flex-row justify-items-start items-center">
            <h2 class="w-20 text-md">Formats:</h2>
            <ButtonGroup class="space-x-px">
                <Button pill color="green" onclick={() => (pdf = !pdf)}>
                    {#if pdf}
                        <CheckOutline />
                    {/if}.pdf
                </Button>
                <Button pill color="purple" onclick={() => (rmdoc = !rmdoc)}>
                    {#if rmdoc}
                        <CheckOutline />
                    {/if}.rmdoc
                </Button>
            </ButtonGroup>
        </div>
        <div class="flex flex-row justify-items-start items-center mt-3">
            <h2 class="w-20 text-md">Location:</h2>
            <Button pill onclick={selectDirectory}>Choose a folder</Button>
            <h2 class="text-md ml-2">{location || "No folder selected."}</h2>
        </div>
        {#if items.length > 0}
            <h1 class="mb-2 mt-4 text-lg font-bold">
                Following items will be exported:
            </h1>
            <Listgroup {items} let:item active={false}>
                <div class="flex flex-row justify-start items-center w-full">
                    <FileLinesSolid class="mr-1" size="lg" />
                    <P size="xl">{item.DisplayPath}</P>
                </div>
            </Listgroup>
        {/if}
    </main>
    <div class="fixed bottom-7 right-10">
        <Button
            disabled={!location || (!pdf && !rmdoc)}
            pill
            size="xl"
            onclick={onProceed}>Proceed</Button
        >
    </div>
</div>
