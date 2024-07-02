package main

import (
    "encoding/binary"
    "fmt"
    "os"
)

/*
** The byte offsets on a file page for various variables. **

** MySQL-4.0.14 space id the page belongs to (== 0) but in later
versions the 'new' checksum of the page
constexpr uint32_t FIL_PAGE_SPACE_OR_CHKSUM = 0;

** page offset inside space
constexpr uint32_t FIL_PAGE_OFFSET = 4;

** if there is a 'natural' predecessor of the page, its offset.
Otherwise FIL_NULL. This field is not set on BLOB pages, which are stored as a
singly-linked list. See also FIL_PAGE_NEXT.
constexpr uint32_t FIL_PAGE_PREV = 8;

** On page 0 of the tablespace, this is the server version ID
constexpr uint32_t FIL_PAGE_SRV_VERSION = 8;

** if there is a 'natural' successor of the page, its offset. Otherwise
FIL_NULL. B-tree index pages(FIL_PAGE_TYPE contains FIL_PAGE_INDEX) on the
same PAGE_LEVEL are maintained as a doubly linked list via FIL_PAGE_PREV and
FIL_PAGE_NEXT in the collation order of the smallest user record on each
page.
constexpr uint32_t FIL_PAGE_NEXT = 12;

** On page 0 of the tablespace, this is the server version ID
constexpr uint32_t FIL_PAGE_SPACE_VERSION = 12;

** lsn of the end of the newest modification log record to the page
constexpr uint32_t FIL_PAGE_LSN = 16;

** file page type: FIL_PAGE_INDEX,..., 2 bytes. The contents of this field
can only be trusted in the following case: if the page is an uncompressed
B-tree index page, then it is guaranteed that the value is FIL_PAGE_INDEX.
The opposite does not hold.

In tablespaces created by MySQL/InnoDB 5.1.7 or later, the contents of this
field is valid for all uncompressed pages.
constexpr uint32_t FIL_PAGE_TYPE = 24;

** this is only defined for the first page of the system tablespace: the file
has been flushed to disk at least up to this LSN. For FIL_PAGE_COMPRESSED
pages, we store the compressed page control information in these 8 bytes.
constexpr uint32_t FIL_PAGE_FILE_FLUSH_LSN = 26;

** If page type is FIL_PAGE_COMPRESSED then the 8 bytes starting at
FIL_PAGE_FILE_FLUSH_LSN are broken down as follows:

** Control information version format (u8)
constexpr uint32_t FIL_PAGE_VERSION = FIL_PAGE_FILE_FLUSH_LSN;

** Compression algorithm (u8)
constexpr uint32_t FIL_PAGE_ALGORITHM_V1 = FIL_PAGE_VERSION + 1;

** Original page type (u16)
constexpr uint32_t FIL_PAGE_ORIGINAL_TYPE_V1 = FIL_PAGE_ALGORITHM_V1 + 1;

** Original data size in bytes (u16)
constexpr uint32_t FIL_PAGE_ORIGINAL_SIZE_V1 = FIL_PAGE_ORIGINAL_TYPE_V1 + 2;

** Size after compression (u16)
constexpr uint32_t FIL_PAGE_COMPRESS_SIZE_V1 = FIL_PAGE_ORIGINAL_SIZE_V1 + 2;

** This overloads FIL_PAGE_FILE_FLUSH_LSN for RTREE Split Sequence Number
constexpr uint32_t FIL_RTREE_SPLIT_SEQ_NUM = FIL_PAGE_FILE_FLUSH_LSN;

** starting from 4.1.x this contains the space id of the page
constexpr uint32_t FIL_PAGE_ARCH_LOG_NO_OR_SPACE_ID = 34;

** alias for space id
constexpr uint32_t FIL_PAGE_SPACE_ID = FIL_PAGE_ARCH_LOG_NO_OR_SPACE_ID;
*/

const PageSize = 16 * 1024

func parseFile(fileName, pageTypeLang string) {
    fileInfo, err := os.Stat(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error getting file info: %v\n", err)
        return
    }

    fileSize := fileInfo.Size()
    if fileSize%PageSize != 0 {
        fmt.Fprintf(os.Stderr, "File size is not a multiple of PAGE_SIZE\n")
        return
    }

    pageCount := fileSize / PageSize
    fmt.Printf("\n- InnoDB tablespace file: %s\n", fileName)
    fmt.Printf("- Pages: %d\n\n", pageCount)

    file, err := os.Open(fileName)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
        return
    }
    defer file.Close()

    for i := int64(0); i < pageCount; i++ {
        page := make([]byte, PageSize)
        _, err := file.Read(page)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading page: %v\n", err)
            return
        }
        analyzePage(page, pageTypeLang)
    }
}

/* 
- fil0types.h
- fil0fil.h 
start of the data on the page
FIL_PAGE_DATA 38U (38byte)
*/
func analyzePage(page []byte, pageTypeLang string) {
    pageTypeMapKr := initPageTypeMapKr()
    pageTypeMap := initPageTypeMap()

    pageType := binary.BigEndian.Uint16(page[24:26])
    convertPageType := fmt.Sprintf("0x%04x", pageType)

    selectedMap := selectMap(pageTypeLang, pageTypeMap, pageTypeMapKr)

    pageTypeStr, exists := selectedMap[convertPageType]
    if !exists {
        pageTypeStr = "Unknown (" + convertPageType + ")" // 맵에 없는 페이지 유형일 경우
    }

    pageNum := int32(binary.BigEndian.Uint32(page[4:8]))   // FIL_PAGE_OFFSET
    pagePrev := int32(binary.BigEndian.Uint32(page[8:12])) // FIL_PAGE_PREV
    pageNext := int32(binary.BigEndian.Uint32(page[12:16])) // FIL_PAGE_NEXT
    recordCount := binary.BigEndian.Uint16(page[54:56]) // 
    treeLevel := binary.BigEndian.Uint16(page[64:66])
    indexId := binary.BigEndian.Uint64(page[66:74])
    //RecordType := binary.BigEndian.Uint16(page[95:97]) & 0x7

    fmt.Printf("page_type = %-40s page_num = %-8d page_prev = %-8d page_next = %-8d record_count = %-8d tree_level = %-8d index_id = %-20d\n",
        pageTypeStr+"("+fmt.Sprintf("0x%04x", pageType)+")", pageNum, pagePrev, pageNext, recordCount, treeLevel, indexId)
}

func initPageTypeMapKr() map[string]string {
    return map[string]string{
        "0x0000": "새로 할당된 페이지", 
        "0x0001": "미사용 페이지", 
        "0x0002": "언두 로그 페이지", 
        "0x0003": "인덱스 노드", 
        "0x0004": "체인지 버퍼 유휴 목록", 
        "0x0005": "체인지 버퍼 비트맵", 
        "0x0006": "시스템 페이지", 
        "0x0007": "트랜잭션 시스템 데이터", 
        "0x0008": "테이블 스페이스 헤더", 
        "0x0009": "익스텐트 디스크립터 페이지", 
        "0x000A": "압축되지 않은 BLOB 페이지", 
        "0x000B": "첫 번째 압축된 BLOB 페이지", 
        "0x000C": "이후 압축된 BLOB 페이지", 
        "0x000D": "알 수 없는 페이지 타입", 
        "0x000E": "압축된 페이지", 
        "0x000F": "암호화된 페이지", 
        "0x0010": "압축 및 암호화된 페이지", 
        "0x0011": "암호화된 R-tree 페이지", 
        "0x0012": "압축되지 않은 SDI BLOB 페이지", 
        "0x0013": "압축된 SDI BLOB 페이지", 
        "0x0014": "레거시 더블라이트 버퍼 페이지", 
        "0x0015": "롤백 세그먼트 배열 페이지", 
        "0x0016": "압축되지 않은 LOB의 인덱스 페이지", 
        "0x0017": "압축되지 않은 LOB의 데이터 페이지", 
        "0x0018": "압축되지 않은 LOB의 첫 번째 페이지", 
        "0x0019": "압축된 LOB의 첫 번째 페이지", 
        "0x001A": "압축된 LOB의 데이터 페이지", 
        "0x001B": "압축된 LOB의 인덱스 페이지", 
        "0x001C": "압축된 LOB의 조각 페이지", 
        "0x001D": "조각 페이지의 인덱스 페이지 (압축된 LOB)", 
        "0x45bf": "인덱스 페이지", 
        "0x45be": "R-tree 노드", 
        "0x45bd": "테이블스페이스 SDI 인덱스 페이지", 
    }
}

func initPageTypeMap() map[string]string {
    return map[string]string{
        "0x0000": "Freshly allocated page", 
        "0x0001": "Unused page", 
        "0x0002": "Undo log page", 
        "0x0003": "Index node", 
        "0x0004": "Insert buffer free list", 
        "0x0005": "Insert buffer bitmap", 
        "0x0006": "System page", 
        "0x0007": "Transaction system data", 
        "0x0008": "File space header", 
        "0x0009": "Extent descriptor page", 
        "0x000A": "Uncompressed BLOB page", 
        "0x000B": "First compressed BLOB page", 
        "0x000C": "Subsequent compressed BLOB page", 
        "0x000D": "Unknown page", 
        "0x000E": "Compressed page", 
        "0x000F": "Encrypted page", 
        "0x0010": "Compressed and Encrypted page", 
        "0x0011": "Encrypted R-tree page", 
        "0x0012": "Uncompressed SDI BLOB page", 
        "0x0013": "Compressed SDI BLOB page", 
        "0x0014": "Legacy doublewrite buffer page", 
        "0x0015": "Rollback Segment Array page", 
        "0x0016": "Index pages of uncompressed LOB", 
        "0x0017": "Data pages of uncompressed LOB", 
        "0x0018": "The first page of an uncompressed LOB", 
        "0x0019": "The first page of a compressed LOB", 
        "0x001A": "Data pages of compressed LOB", 
        "0x001B": "Index pages of compressed LOB", 
        "0x001C": "Fragment pages of compressed LOB", 
        "0x001D": "Index pages of fragment pages (compressed LOB)", 
        "0x45bf": "B-tree node", 
        "0x45be": "R-tree node", 
        "0x45bd": "Tablespace SDI Index page", 
    }
}

func selectMap(pageTypeLang string, pageTypeMap, pageTypeMapKr map[string]string) map[string]string {
    if pageTypeLang == "1" {
        return pageTypeMapKr
    }
    return pageTypeMap
}

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <filename> [pageTypeLang]\n", os.Args[0])
        return
    }

    fileName := os.Args[1]
    pageTypeLang := "0"
    if len(os.Args) >= 3 {
        pageTypeLang = os.Args[2]
    }

    parseFile(fileName, pageTypeLang)
}
