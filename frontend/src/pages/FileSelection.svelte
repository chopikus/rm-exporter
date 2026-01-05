<script lang="ts">
    import {
        Button,
        Checkbox,
        Listgroup,
        Navbar,
        P,
        ToolbarButton,
        Tooltip,
    } from "flowbite-svelte";
    import {
        ArrowUpOutline,
        FileLinesSolid,
        FolderSolid,
    } from "flowbite-svelte-icons";
    import {
        GetFolder,
        GetFolderSelection,
        GetItemSelection,
        OnItemSelect,
        GetCheckedFilesCount,
    } from "../../wailsjs/go/main/App";
    import { push } from "svelte-spa-router";
    import { backend } from "../../wailsjs/go/models";
    import FileSelectionHeader from "./FileSelectionHeader.svelte";
    import FileSelectionList from "./FileSelectionList.svelte";
    type DocInfo = backend.DocInfo;

    let folderId = $state("");
    let path: string[] = $state([]);
    let items: DocInfo[] = $state([]);

    // Checked stores checkbox value for every element of the folder
    let checked: { [key: string]: number } = $state({});
    // defined in go code
    const UNSELECTED = 0,
        INDETERMINATE = 1,
        SELECTED = 2;

    let export_disabled = $state(true);
    GetCheckedFilesCount().then((count: number) => {
        export_disabled = count === 0;
    });

    // onIdUpdate
    $effect(() => {
        GetFolder(folderId).then((result) => {
            items = result;
        });
        GetFolderSelection(folderId).then((result) => {
            checked = {};
            for (const item of result) {
                checked[item.Id] = item.Status;
            }
        });
    });

    const onBack = () => {
        if (path.length > 0) {
            folderId = path[path.length - 1];
            path.pop();
        } else {
            folderId = "";
        }
    };

    const onItemClick = (item: DocInfo) => {
        if (item.IsFolder) {
            path.push(folderId);
            folderId = item.Id;
        }
    };

    const onExportClick = () => {
        //storeCheckedFiles();
        push("/export-confirmation");
    };

    const isItemChecked = (id: string) => {
        return checked[id] === SELECTED;
    };

    const isItemIndeterminate = (id: string) => {
        return checked[id] === INDETERMINATE;
    };

    const itemCheckUpdate = (id: string, value: boolean | undefined) => {
        let select;
        if (value) {
            checked[id] = SELECTED;
            select = true;
        } else {
            checked[id] = UNSELECTED;
            select = false;
        }
        OnItemSelect(id, select).then(() => {
            if (id === folderId) {
                GetFolderSelection(folderId).then((result) => {
                    checked = {};
                    for (const item of result) {
                        checked[item.Id] = item.Status;
                    }
                });
            } else {
                GetItemSelection(folderId).then((result) => {
                    checked[result.Id] = result.Status;
                });
            }
            GetCheckedFilesCount().then((count: number) => {
                export_disabled = count === 0;
            });
        });
    };
</script>

<div style="height: fit-content;">
    <FileSelectionHeader
        id={folderId}
        {path}
        {onBack}
        {isItemChecked}
        {isItemIndeterminate}
        {itemCheckUpdate}
    />
    <main class="pl-10 pr-10 pt-3 pb-3">
        <FileSelectionList
            {items}
            {isItemChecked}
            {isItemIndeterminate}
            {itemCheckUpdate}
            {onItemClick}
        />
    </main>
    <div class="fixed bottom-7 right-10">
        <Button
            pill
            size="xl"
            disabled={export_disabled}
            onclick={onExportClick}>Export</Button
        >
    </div>
</div>
