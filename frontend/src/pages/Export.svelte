<script lang="ts">
    import {
        Alert,
        Button,
        Listgroup,
        Navbar,
        P,
        Spinner,
    } from "flowbite-svelte";
    import {
        InitExport,
        Export,
        GetCheckedFiles,
        GetExportOptions,
    } from "../../wailsjs/go/main/App.js";
    import {
        CheckOutline,
        ExclamationCircleOutline,
        FileLinesSolid,
        InfoCircleSolid,
    } from "flowbite-svelte-icons";
    import { EventsOn } from "../../wailsjs/runtime/runtime.js";
    import { backend } from "../../wailsjs/go/models.js";
    type DocInfo = backend.DocInfo;

    let exportItems: DocInfo[] = $state([]);
    let exportItemState: { [key: string]: string } = $state({});
    let exportOptions: backend.RmExportOptions = $state({});

    let errorMessage: string = $state("hello");
    let showError: boolean = $state(false);
    let failed: boolean = $state(false);

    GetCheckedFiles().then((result: DocInfo[]) => {
        exportItems = result;
    });

    GetExportOptions().then((result: backend.RmExportOptions) => {
        exportOptions = result;
    });

    let formats: string = $derived.by(() => {
        let result = "";
        if (exportOptions.Pdf) {
            result += "pdf";
        }
        if (exportOptions.Rmdoc) {
            if (result.length != 0) {
                result += ",";
            }
            result += "rmdoc";
        }
        return result;
    });

    let dirStructs: string = $derived.by(() => {
        if (exportOptions.PreserveDirStruct) {
            return "Preserving directory structure.";
        } else {
            return "Downloading everything into the same folder.";
        }
    });

    let timeoutSeconds: number = $derived.by(() => {
        return exportOptions.TimeoutSeconds;
    });

    InitExport().then(() => {
        Export();
    });

    const onRetry = () => {
        failed = false;
        Export();
    };

    const finishedAllItems = $derived(() => {
        return (
            Object.keys(exportItemState).filter(
                (id: string) => exportItemState[id] == "finished",
            ).length == exportItems.length
        );
    });

    EventsOn("started", (id: string) => {
        exportItemState[id] = "started";
    });

    EventsOn("finished", (id: string) => {
        exportItemState[id] = "finished";
    });

    EventsOn("failed", (id: string, msg: string) => {
        exportItemState[id] = "failed";
        showError = true;
        errorMessage = msg;
        failed = true;
    });
</script>

<div style="height: fit-content;">
    {#key showError}
        {#if showError}
            <Alert color="red" dismissable={true}>
                <span slot="icon">
                    <InfoCircleSolid class="w-5 h-5" />
                    <span class="sr-only">Info</span>
                </span>
                <p class="font-large">Couldn't export a file!</p>
                <p class="font-medium">{errorMessage}</p>
            </Alert>
        {/if}
    {/key}

    <Navbar color="blue" class="sticky top-0">
        <h1 class="font-bold m-auto">
            {finishedAllItems() ? "Success!" : "Export"}
        </h1>
    </Navbar>

    <main class="pr-7 pl-7 mt-3 w-full">
        <h2 class="text-md">Formats: {formats}</h2>
        <h2 class="text-md">Location: {exportOptions["Location"]}</h2>
        <h2 class="text-md">
            Download timeout per note: {(timeoutSeconds / 60).toFixed(1)} min
        </h2>
        <h2 class="text-md mb-3">{dirStructs}</h2>

        {#if exportItems.length > 0}
            <Listgroup items={exportItems} let:item active={false}>
                <div class="flex flex-row justify-start items-center w-full">
                    <FileLinesSolid class="mr-1" size="lg" />
                    <P size="xl">{item.DisplayPath}</P>
                    {#if exportItemState[item.Id] === "started"}
                        <Spinner class="ml-auto" />
                    {:else if exportItemState[item.Id] === "finished"}
                        <CheckOutline class="ml-auto" color="green" />
                    {:else if exportItemState[item.Id] === "failed"}
                        <ExclamationCircleOutline class="ml-auto" color="red" />
                    {/if}
                </div>
            </Listgroup>
        {/if}
    </main>

    <div class="fixed bottom-7 right-10">
        <Button
            class={!failed ? "invisible" : ""}
            pill
            size="xl"
            onclick={onRetry}>Retry</Button
        >
    </div>
</div>
